package com.DExam.User_Service.service;

import com.DExam.User_Service.database.UserRepository;
import com.DExam.User_Service.model.User;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.Mock;
import org.mockito.MockitoAnnotations;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;

import java.util.Optional;

import static org.assertj.core.api.AssertionsForClassTypes.assertThat;
import static org.mockito.Mockito.when;

class UserServiceTest {
    @Mock
    private UserRepository userRepository;
    private AutoCloseable autoCloseable;
    private UserService underTest;

    @BeforeEach
    void setUp() {
        autoCloseable = MockitoAnnotations.openMocks(this);
        underTest = new UserService(userRepository, new BCryptPasswordEncoder());
    }

    @AfterEach
    void tearDown() throws Exception {
        autoCloseable.close();
    }

    @Test
    void get() {
        String[] emails = new String[]{"eyad@gmail.com", "marwan@gmail.com",
                "ammar@gmail.com", "bodda@gmail.com"};
        for (int i = 0; i < emails.length; i++) {
            User user = new User();
            user.setEmail(emails[i]);
            if (i == emails.length - 1)
                when(userRepository.findByEmail(emails[i])).thenReturn(Optional.empty());
            else
                when(userRepository.findByEmail(emails[i])).thenReturn(Optional.of(user));
            Optional<User> exist = Optional.empty();
            try {
                exist = Optional.ofNullable(underTest.get(emails[i]));
                assertThat(exist).isNotEqualTo(Optional.empty());
                System.out.println("found");
            } catch (Exception exception) {
                assertThat(exist).isEqualTo(Optional.empty());
                System.out.println("not found");
            }
        }
    }


    @Test
    void exists() {
        //testing emails
        String[] emails = new String[]{"eyad@gmail.com", "marwan@gmail.com",
                "ammar@gmail.com", "bodda@gmail.com"};
        for (int i = 0; i < emails.length; i++) {
            User user = new User();
            user.setEmail(emails[i]);
            if (i == emails.length - 1)
                when(userRepository.findByEmail(emails[i])).thenReturn(Optional.empty());
            else
                when(userRepository.findByEmail(emails[i])).thenReturn(Optional.of(user));
            Optional<User> exist = Optional.empty();
            try {
                exist = Optional.ofNullable(underTest.get(emails[i]));
                assertThat(exist).isNotEqualTo(Optional.empty());
                System.out.println("found");
            } catch (Exception exception) {
                assertThat(exist).isEqualTo(Optional.empty());
                System.out.println("not found");
            }
        }

        String[] ids = new String[]{"0132156", "9784654",
                "9876541", "16531320231"};
        for (int i = 0; i < ids.length; i++) {
            User user = new User();
            user.setNationalID(ids[i]);
            if (i == ids.length - 1)
                when(userRepository.findByEmail(ids[i])).thenReturn(Optional.empty());
            else
                when(userRepository.findByEmail(ids[i])).thenReturn(Optional.of(user));
            Optional<User> exist = Optional.empty();
            try {
                exist = Optional.ofNullable(underTest.get(ids[i]));
                assertThat(exist).isNotEqualTo(Optional.empty());
                System.out.println("found");
            } catch (Exception exception) {
                assertThat(exist).isEqualTo(Optional.empty());
                System.out.println("not found");
            }
        }
    }
}