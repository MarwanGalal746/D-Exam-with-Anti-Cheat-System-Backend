package com.DExam.User_Service.model;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class ResetPassRequest {
    private String email;
    private String currentPassword;
    private String newPassword;
}
