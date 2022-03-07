package com.DExam.User_Service.controllers;

import com.DExam.User_Service.modules.User;
import org.springframework.web.bind.annotation.*;

import java.util.Map;

@RestController
@RequestMapping("/api/users")
public class UserController {


    @GetMapping("/get")
    public User get(@RequestParam(value = "id") int id){

        return null;
    }
    @PostMapping("/create")
    public int post(@RequestBody Map<String, Object> userInfo){

        return 0;
    }
    @DeleteMapping("/delete")
    public boolean delete(){

        return true;
    }
    @PutMapping("/update")
    public User put(){

        return null;
    }
}
