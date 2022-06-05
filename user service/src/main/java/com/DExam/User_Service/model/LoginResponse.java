package com.DExam.User_Service.model;

import com.DExam.User_Service.domain.User;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor @NoArgsConstructor
public class LoginResponse {
    private String token;
    private User user;
}
