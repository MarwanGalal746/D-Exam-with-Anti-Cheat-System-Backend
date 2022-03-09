package com.DExam.User_Service.resources.controllers;

import com.DExam.User_Service.resources.modules.Role;
import com.DExam.User_Service.resources.modules.User;
import com.DExam.User_Service.resources.services.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/api/users")
public class UserController {
    private final UserService userService;

    @Autowired
    public UserController(UserService userService) {
        this.userService = userService;
    }

    @GetMapping("/get")
    public User get(@RequestParam long id){
        return userService.get(id);
    }

    @PostMapping("/create")
    public long post(@RequestBody User userInfo){
        return userService.add(userInfo);
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

}
