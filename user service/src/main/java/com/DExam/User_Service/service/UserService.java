package com.DExam.User_Service.service;

import com.DExam.User_Service.database.UserRepository;
import com.DExam.User_Service.exception.EmailExistException;
import com.DExam.User_Service.exception.NationalIDException;
import com.DExam.User_Service.exception.UserNotFoundException;
import com.DExam.User_Service.model.User;
import lombok.RequiredArgsConstructor;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.stereotype.Service;

import java.util.ArrayList;

@Service
@RequiredArgsConstructor
public class UserService implements UserDetailsService {

    private final UserRepository userRepository;
    private final BCryptPasswordEncoder bCryptPasswordEncoder;

    public User get(long id) {
        return userRepository.findById(id)
                .orElseThrow(UserNotFoundException::new);
    }

    public User get(String email) {
        return userRepository.findByEmail(email)
                .orElseThrow(UserNotFoundException::new);
    }

    public long add(User newUser) {
        newUser.setPassword(bCryptPasswordEncoder.encode(newUser.getPassword()));
        userRepository.save(newUser);
        return newUser.getId();
    }

    public void userExistByEmail(String email){
        if (userRepository.findByEmail(email).orElse(null) != null)
            throw new EmailExistException();
    }

    public void userExistByNationalID(String nationalID){
         if (userRepository.findByNationalID(nationalID).orElse(null) != null)
            throw new NationalIDException();
    }

    @Override
    public UserDetails loadUserByUsername(String email) throws UsernameNotFoundException {
        User user = userRepository.findByEmail(email)
                .orElseThrow(UserNotFoundException::new);

        return new org.springframework.security.core.userdetails.User(user.getEmail(),user.getPassword(),new ArrayList<>());
    }

    public void resetPassword(String email, String newPassword) {
        newPassword = bCryptPasswordEncoder.encode(newPassword);
        userRepository.updatePassword(email, newPassword);
    }
}
