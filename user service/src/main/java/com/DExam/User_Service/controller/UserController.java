package com.DExam.User_Service.controller;

import com.DExam.User_Service.config.JwtManager;
import com.DExam.User_Service.exception.InvalidEmailPasswordException;
import com.DExam.User_Service.model.*;
import com.DExam.User_Service.service.IUserService;
import com.DExam.User_Service.service.UserService;
import com.DExam.User_Service.utility.CodeGenerator;
import com.DExam.User_Service.utility.CustomResponse;
import lombok.AllArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.web.bind.annotation.*;

@RestController
@CrossOrigin(origins = "*", allowedHeaders = "*")
@RequestMapping("/api/users")
@AllArgsConstructor
@Slf4j
public class UserController {

    private final IUserService userService;
    private final JwtManager jwtManager;
    private final AuthenticationManager authenticationManager;
    private EmailController emailController;

    @PostMapping("/verify")
    public ResponseEntity<?> verify(@RequestBody User user){
        userService.userExistByEmail(user.getEmail());
        userService.userExistByNationalID(user.getNationalID());
        String verificationCode = CodeGenerator.generateCode();
        MailForm mailForm = new MailForm(user.getEmail(),"EMAIL VERIFICATION",CustomResponse.EMAIL_VERIFICATION + verificationCode);
        emailController.send(mailForm);
        log.info("a verification email has been sent to this email " + user.getEmail());
        return new ResponseEntity<>(new CustomResponse().setMessage(verificationCode).setStatus(HttpStatus.OK),HttpStatus.OK);
    }

    @PostMapping("/register")
    public ResponseEntity<?> register(@RequestBody User user){
        long userID = userService.save(user);
        log.info("a new user has been added with id " + userID);
        return new ResponseEntity<>(new CustomResponse().setMessage(String.valueOf(userID)).setStatus(HttpStatus.CREATED),HttpStatus.CREATED);
    }

    @PutMapping("/update")
    public ResponseEntity<?> update(@RequestBody UpdateUserRequest request, @RequestHeader (name="Authorization") String token){
        token = token.split(" ")[1];
        boolean isValid  = jwtManager.validateToken(token, request.getOldUser());

        if(!isValid){
            log.error("the token is not valid for the user with email " + request.getOldUser().getId());
            return new ResponseEntity<>(new CustomResponse().setMessage(CustomResponse.INVALID_TOKEN).setStatus(HttpStatus.NOT_ACCEPTABLE),HttpStatus.NOT_ACCEPTABLE);
        }
        userService.userExistByEmail(request.getNewUser().getEmail());
        userService.save(request.getNewUser());
        log.info("the credentials of the user with id " + request.getNewUser().getId() + " have been updated" );
        String newToken = jwtManager.generateToken(request.getNewUser().getEmail());
        return new ResponseEntity<>(new CustomResponse().setMessage(newToken).setStatus(HttpStatus.OK),HttpStatus.OK);
    }

    @PostMapping("/login")
    public LoginResponse login(@RequestBody UserCredentials userCredentials) {
        try {
            authenticationManager.authenticate(
                    new UsernamePasswordAuthenticationToken(
                            userCredentials.getEmail(), userCredentials.getPassword()));
        } catch (Exception exception){
            log.error("email or password or both of the user with email " + userCredentials.getEmail() + " are not valid" );
            throw new InvalidEmailPasswordException();
        }

        String accessToken = jwtManager.generateToken(userCredentials.getEmail());
        log.info("user with email " + userCredentials.getEmail() + " has signed in successfully" );
        return new LoginResponse(accessToken, userService.get(userCredentials.getEmail()));
    }

    @PutMapping("/reset")
    public ResponseEntity<?> reset(@RequestBody UserCredentials userCredentials){
        userService.resetPassword(userCredentials.getEmail(), userCredentials.getPassword());
        log.info("password of the user with email " + userCredentials.getEmail() + " has been updated successfully");
        return new ResponseEntity<>(new CustomResponse().setMessage(CustomResponse.PASS_UPDATED).setStatus(HttpStatus.OK),HttpStatus.OK);

    }
}
