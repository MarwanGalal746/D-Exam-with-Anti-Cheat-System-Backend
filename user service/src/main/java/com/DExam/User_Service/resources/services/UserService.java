package com.DExam.User_Service.resources.services;

import com.DExam.User_Service.resources.database.UserRepository;
import com.DExam.User_Service.resources.modules.User;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.stereotype.Service;

@Service
public class UserService {

    private final UserRepository userRepository;

    @Autowired
    private BCryptPasswordEncoder passwordEncoder;

    @Autowired
    public UserService(UserRepository userRepository) {
        this.userRepository = userRepository;
    }

    public User get(long id) {
        return userRepository.findById(id).orElse(null);
    }

    public long add(User newUser) {
        if (userRepository.findByEmail(newUser.getEmail()) != null)
            return -1;
        else if (userRepository.findByNationalID(newUser.getNationalID()) != null)
            return -2;
        newUser.setPassword(passwordEncoder.encode(newUser.getPassword()));
        userRepository.save(newUser);
        return newUser.getId();
    }

    public boolean delete(long id) {
        userRepository.deleteById(id);
        return true;
    }


}
