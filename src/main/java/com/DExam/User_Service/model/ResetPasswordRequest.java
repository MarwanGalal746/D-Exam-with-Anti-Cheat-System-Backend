package com.DExam.User_Service.model;

import lombok.AllArgsConstructor;
import lombok.Data;

@Data
@AllArgsConstructor
public class ResetPasswordRequest {
    private String email;
    private String oldpassword;
    private String newpassword;
}
