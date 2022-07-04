package com.DExam.User_Service.service;

import com.DExam.User_Service.domain.User;

public interface IUserService {
    User get(String email);
    long save(User user);
    void userExistByEmail(String email);
    void userExistByNationalID(String nationalID);
    void resetPassword(String email, String password);
    boolean activateUser(String email, String password);
    boolean isUserActive(String email);
}
