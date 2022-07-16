package com.DExam.User_Service.service;

import com.DExam.User_Service.domain.User;
import com.DExam.User_Service.model.CourseStudentsInfo;

import java.util.ArrayList;

public interface IUserService {
    User get(String email);
    long save(User user);
    void userExistByEmail(String email);
    void userExistByNationalID(String nationalID);
    void updatePassword(String email, String oldPassword, String newPassword);
    long activateUser(String email, String password);
    boolean isUserActive(String email);

    ArrayList<CourseStudentsInfo> getUsers(ArrayList<Long> userIDs);
}
