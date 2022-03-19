package com.DExam.User_Service.controller;

import com.DExam.User_Service.exception.InvalidEmailPasswordException;
import com.DExam.User_Service.model.MailForm;
import com.DExam.User_Service.service.EmailService;
import com.DExam.User_Service.model.AuthenticationRequest;
import com.DExam.User_Service.model.User;
import com.DExam.User_Service.service.UserService;
import com.DExam.User_Service.utility.CodeGenerator;
import com.DExam.User_Service.utility.CustomResponse;
import com.DExam.User_Service.config.JwtManager;
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
        userService.exists(user);
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
    public boolean update(@RequestBody User user){
        userService.add(user);
        return true;
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
}
