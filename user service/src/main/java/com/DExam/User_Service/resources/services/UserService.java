package com.DExam.User_Service.resources.services;

import com.DExam.User_Service.resources.database.UserRepository;
import com.DExam.User_Service.resources.modules.User;
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
                .orElseThrow(()->new UsernameNotFoundException("USER DOES NOT EXIST"));
    }

    public User get(String email) {
        return userRepository.findByEmail(email)
                .orElseThrow(()->new UsernameNotFoundException("USER DOES NOT EXIST"));
    }

    public long add(User newUser) {
        if (userRepository.findByEmail(newUser.getEmail()) != null)
            return -1;
        else if (userRepository.findByNationalID(newUser.getNationalID()) != null)
            return -2;

        newUser.setPassword(bCryptPasswordEncoder.encode(newUser.getPassword()));
        userRepository.save(newUser);
        return newUser.getId();
    }

    public boolean delete(long id) {
        userRepository.deleteById(id);
        return true;
    }

    @Override
    public UserDetails loadUserByUsername(String email) throws UsernameNotFoundException {
        User user = userRepository.findByEmail(email)
                .orElseThrow(()->new UsernameNotFoundException("USER DOES NOT EXIST"));

        /**
            Need to pass roles here when it's done to the array
         */
        Collection<SimpleGrantedAuthority> authorities = new ArrayList<>();
        return new org.springframework.security.core.userdetails.User(user.getEmail(),user.getPassword(),authorities);
    }
}
