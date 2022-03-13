package com.DExam.User_Service.resources.controllers;

import com.DExam.User_Service.resources.entity.AuthenticationRequest;
import com.DExam.User_Service.resources.utility.Errors;
import com.DExam.User_Service.resources.entity.User;
import com.DExam.User_Service.resources.services.UserService;
import com.DExam.User_Service.resources.utility.JwtManager;
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

    @GetMapping("/get")
    public User get(@RequestParam long id){
        return userService.get(id);
    }

    @PostMapping("/register")
    public ResponseEntity<?> register(@RequestBody User user){
        long userId = userService.add(user);
        if (userId == -1)
            return new ResponseEntity<>(customResponse("Error", Errors.EMAIL_USED), HttpStatus.BAD_REQUEST);
        else if (userId == -2)
            return new ResponseEntity<>(customResponse("Error", Errors.NATIONAL_ID_USED), HttpStatus.BAD_REQUEST);
        return new ResponseEntity<>(customResponse("Id", userId), HttpStatus.OK);
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
        }catch (Exception exception){
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
