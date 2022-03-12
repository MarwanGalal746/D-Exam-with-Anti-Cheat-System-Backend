package com.DExam.User_Service.resources.controllers;

import com.DExam.User_Service.resources.modules.Errors;
import com.DExam.User_Service.resources.modules.User;
import com.DExam.User_Service.resources.services.UserService;
import lombok.AllArgsConstructor;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.HashMap;

@RestController
@RequestMapping("/api/users")
@AllArgsConstructor
public class UserController {
    private final UserService userService;

    @GetMapping("/getbyid")
    public User get(@RequestParam long id){
        return userService.get(id);
    }

    @GetMapping("/getbyemail")
    public User get(@RequestParam String email){
        return userService.get(email);
    }

    @PostMapping("/create")
    public ResponseEntity<?> post(@RequestBody User userInfo){
        long userId = userService.add(userInfo);
        if (userId == -1)
            return new ResponseEntity<>(customResponse("Error", Errors.EMAIL_USED), HttpStatus.BAD_REQUEST);
        else if (userId == -2)
            return new ResponseEntity<>(customResponse("Error", Errors.NATIONAL_ID_USED), HttpStatus.BAD_REQUEST);
        return new ResponseEntity<>(customResponse("Id", userId), HttpStatus.OK);
    }

    @DeleteMapping("/delete")
    public boolean delete(@RequestParam long id){
        return userService.delete(id);
    }

    @PutMapping("/update")
    public long put(@RequestBody User userInfo){
        userService.delete(userInfo.getId());
        return userService.add(userInfo);
    }

    private HashMap<String, Object> customResponse(String key, Object value) {
        HashMap<String, Object> response = new HashMap();
        response.put(key, value);
        return response;
    }
}
