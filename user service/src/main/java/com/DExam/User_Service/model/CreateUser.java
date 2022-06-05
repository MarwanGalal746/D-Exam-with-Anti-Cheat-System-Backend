package com.DExam.User_Service.model;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter
@Setter
@NoArgsConstructor
@AllArgsConstructor
public class CreateUser {
    private String name;
    private String email;
    private String nationalID;
    private String password;
    private String img;
}
