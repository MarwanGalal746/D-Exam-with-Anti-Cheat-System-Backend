package com.DExam.User_Service.service;

import com.DExam.User_Service.database.UserRepository;
import com.DExam.User_Service.exception.EmailExistException;
import com.DExam.User_Service.exception.NationalIDException;
import com.DExam.User_Service.exception.UserNotFoundException;
import com.DExam.User_Service.model.User;
import lombok.RequiredArgsConstructor;
import org.springframework.security.core.authority.SimpleGrantedAuthority;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.Collection;

@Service
@RequiredArgsConstructor
public class UserService implements UserDetailsService {

    private final UserRepository userRepository;
    private final BCryptPasswordEncoder bCryptPasswordEncoder;

    public User get(long id) {
        return userRepository.findById(id)
                .orElseThrow(()->new UserNotFoundException());
    }

    public User get(String email) {
        return userRepository.findByEmail(email)
                .orElseThrow(()->new UserNotFoundException());
    }

    public long add(User newUser) {
        newUser.setPassword(bCryptPasswordEncoder.encode(newUser.getPassword()));
        userRepository.save(newUser);
        return newUser.getId();
    }

    public void exists(User newUser){
        if (userRepository.findByEmail(newUser.getEmail()).orElse(null) != null)
            throw new EmailExistException();
        else if (userRepository.findByNationalID(newUser.getNationalID()).orElse(null) != null)
            throw new NationalIDException();
    }

    public boolean delete(long id) {
        userRepository.deleteById(id);
        return true;
    }

    @Override
    public UserDetails loadUserByUsername(String email) throws UsernameNotFoundException {
        User user = userRepository.findByEmail(email)
                .orElseThrow(()->new UserNotFoundException());

        return new org.springframework.security.core.userdetails.User(user.getEmail(),user.getPassword(),new ArrayList<>());
    }
}
