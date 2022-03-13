package com.DExam.User_Service.resources.controllers;

import com.DExam.User_Service.resources.services.EmailService;
import com.DExam.User_Service.resources.entity.AuthenticationRequest;
import com.DExam.User_Service.resources.entity.User;
import com.DExam.User_Service.resources.services.UserService;
import com.DExam.User_Service.resources.utility.CodeGenerator;
import com.DExam.User_Service.resources.utility.Errors;
import com.DExam.User_Service.resources.utility.JwtManager;
import com.DExam.User_Service.resources.utility.Templates;
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
    private final EmailService emailSender;

    @GetMapping("/get")
    public User get(@RequestParam long id){
        return userService.get(id);
    }

    @PostMapping("/verify")
    public ResponseEntity<?> verify(@RequestBody User user){
        int status = userService.exists(user);
        if (status == -1)
            return new ResponseEntity<>(customResponse("Error", Errors.EMAIL_USED), HttpStatus.BAD_REQUEST);
        else if (status == -2)
            return new ResponseEntity<>(customResponse("Error", Errors.NATIONAL_ID_USED), HttpStatus.BAD_REQUEST);
        else
        {
            String verificationCode = CodeGenerator.generateCode();
            emailSender.send(user.getEmail(),"EMAIL VERIFICATION",
                    Templates.getTemplates().get("EMAIL VERIFICATION") + verificationCode);

            return new ResponseEntity<>(customResponse("Code", verificationCode), HttpStatus.OK);
        }
    }

    @PostMapping("/register")
    public ResponseEntity<?> register(@RequestBody User user){
        long userID = userService.add(user);
        return new ResponseEntity<>(customResponse("Id", userID), HttpStatus.OK);
    }

    @PutMapping("/update")
    public long update(@RequestBody User userInfo){
        userService.delete(userInfo.getId());
        return userService.add(userInfo);
    }

    @PostMapping("/login")
    public Object login(@RequestBody AuthenticationRequest authenticationRequest) throws Exception {
        try {
            authenticationManager.authenticate(
                    new UsernamePasswordAuthenticationToken(
                            authenticationRequest.getEmail(),authenticationRequest.getPassword()));
        } catch (Exception exception){
                throw new Exception("Invalid email or password");
        }
        String accessToken = jwtManager.generateToken(authenticationRequest.getEmail());
        Map<String, Object> userInfo = new HashMap<>();
        userInfo.put("user",userService.get(authenticationRequest.getEmail()));
        userInfo.put("access_token",accessToken);
        return userInfo;
    }

    private HashMap<String, Object> customResponse(String key, Object value) {
        HashMap<String, Object> response = new HashMap();
        response.put(key, value);
        return response;
    }
}
