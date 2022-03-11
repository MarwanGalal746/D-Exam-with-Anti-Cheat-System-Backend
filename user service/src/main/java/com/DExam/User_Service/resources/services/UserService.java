package com.DExam.User_Service.resources.services;

import com.DExam.User_Service.resources.database.UserRepository;
import com.DExam.User_Service.resources.modules.User;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class UserService {
    @Autowired
    private final UserRepository userRepository;

    @Autowired
    public UserService(UserRepository userRepository) {
        this.userRepository = userRepository;
    }

    public User get(long id) {
        return userRepository.findById(id).orElse(null);
    }

    public long add(User newUser) {
        if (userRepository.isEmailExists(newUser.getEmail()) > 0)
            return -1;
        else if (userRepository.isNationalIdExists(newUser.getNationalID()) > 0)
            return -2;
        userRepository.save(newUser);
        return newUser.getId();
    }

    public boolean delete(long id) {
        userRepository.deleteById(id);
        return true;
    }


}
