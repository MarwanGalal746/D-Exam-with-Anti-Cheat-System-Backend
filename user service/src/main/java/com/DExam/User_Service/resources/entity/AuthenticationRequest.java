package com.DExam.User_Service.resources.entity;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor @NoArgsConstructor
public class AuthenticationRequest {
    private String email;
    private String password;
}
