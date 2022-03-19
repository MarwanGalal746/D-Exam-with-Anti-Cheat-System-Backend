package com.DExam.User_Service.database;

import com.DExam.User_Service.model.User;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import java.util.Optional;

import static org.assertj.core.api.AssertionsForClassTypes.assertThat;

@SpringBootTest

class UserRepositoryTest {
    @Autowired
    private UserRepository underTest;

    @Test
    void findByEmailTest() {
        User user = new User();
        String email = "eyaaad@gmail.com";
        user.setEmail(email);
        Optional exist = underTest.findByEmail(email);
        assertThat(exist).isEqualTo(Optional.empty());
    }
}