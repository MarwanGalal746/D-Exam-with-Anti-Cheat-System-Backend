package com.DExam.User_Service.controller;

import com.DExam.User_Service.config.JwtManager;
import com.DExam.User_Service.exception.InvalidEmailPasswordException;
import com.DExam.User_Service.model.*;
import com.DExam.User_Service.service.UserService;
import com.DExam.User_Service.utility.CodeGenerator;
import com.DExam.User_Service.utility.CustomResponse;
import lombok.AllArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.web.bind.annotation.*;

import java.util.HashMap;
import java.util.Map;

@RestController
@RequestMapping("/api/users")
@AllArgsConstructor
public class UserController {

    private final UserService userService;
    private final JwtManager jwtManager;
    private final AuthenticationManager authenticationManager;
    private EmailController emailController;

    @GetMapping("/get")
    public User get(@RequestParam long id){
        return userService.get(id);
    }

    @PostMapping("/verify")
    public ResponseEntity<?> verify(@RequestBody User user){
        userService.userExistByEmail(user.getEmail());
        userService.userExistByNationalID(user.getNationalID());
        String verificationCode = CodeGenerator.generateCode();
        MailForm mailForm = new MailForm(user.getEmail(),"EMAIL VERIFICATION",CustomResponse.EMAIL_VERIFICATION + verificationCode);
        emailController.send(mailForm);
        return new ResponseEntity<>(new CustomResponse().setMessage(verificationCode).setStatus(HttpStatus.OK),HttpStatus.OK);
    }

    @PostMapping("/register")
    public ResponseEntity<?> register(@RequestBody User user){
        long userID = userService.add(user);
        return new ResponseEntity<>(new CustomResponse().setMessage(String.valueOf(userID)).setStatus(HttpStatus.CREATED),HttpStatus.CREATED);
    }

    @PutMapping("/update")
    public ResponseEntity<?> update(@RequestBody UpdateUserRequest request, @RequestHeader (name="Authorization") String token){
        token = token.split(" ")[1];
        boolean isValid  = jwtManager.validateToken(token, request.getOldUser());

        if(!isValid)
            return new ResponseEntity<>(new CustomResponse().setMessage(CustomResponse.INVALID_TOKEN).setStatus(HttpStatus.NOT_ACCEPTABLE),HttpStatus.NOT_ACCEPTABLE);

        userService.userExistByEmail(request.getNewUser().getEmail());
        userService.add(request.getNewUser());
        String newToken = jwtManager.generateToken(request.getNewUser().getEmail());
        return new ResponseEntity<>(new CustomResponse().setMessage(newToken).setStatus(HttpStatus.OK),HttpStatus.OK);
    }

    @PostMapping("/login")
    public Object login(@RequestBody AuthenticationRequest authenticationRequest) {
        try {
            authenticationManager.authenticate(
                    new UsernamePasswordAuthenticationToken(
                            authenticationRequest.getEmail(),authenticationRequest.getPassword()));
        } catch (Exception exception){
                throw new InvalidEmailPasswordException();
        }

        String accessToken = jwtManager.generateToken(authenticationRequest.getEmail());

        Map<String, Object> userInfo = new HashMap<>();
        userInfo.put("user",userService.get(authenticationRequest.getEmail()));
        userInfo.put("access_token",accessToken);

        return userInfo;
    }

    @PutMapping("/reset")
    public ResponseEntity<?> reset(@RequestBody ResetPassRequest resetPassRequest){
        try {
            authenticationManager.authenticate(
                    new UsernamePasswordAuthenticationToken(
                            resetPassRequest.getEmail(),resetPassRequest.getCurrentPassword()));
        } catch (Exception exception){
            throw new InvalidEmailPasswordException();
        }
        userService.resetPassword(resetPassRequest.getEmail(),resetPassRequest.getCurrentPassword());
        return new ResponseEntity<>(new CustomResponse().setMessage(CustomResponse.PASS_UPDATED).setStatus(HttpStatus.OK),HttpStatus.OK);

    }
}
