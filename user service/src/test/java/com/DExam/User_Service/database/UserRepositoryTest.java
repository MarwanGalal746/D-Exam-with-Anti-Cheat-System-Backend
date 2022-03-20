package com.DExam.User_Service.database;

import com.DExam.User_Service.model.User;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import java.util.ArrayList;
import java.util.Optional;

import static org.assertj.core.api.AssertionsForClassTypes.assertThat;

@SpringBootTest

class UserRepositoryTest {
    @Autowired
    private UserRepository underTest;

    @Test
    void findByEmailTest() {
        String[] emails = new String[] { "eyad@gmail.com", "marwan@gmail.com", "ammar@gmail.com", "bodda@gmail.com"};
        ArrayList<User> users = new ArrayList<>();
        for (int i=0 ; i< emails.length-1 ; i++) {
            User user  = new User();
            user.setEmail(emails[i]);
            underTest.save(user);
            users.add(user);
        }

        for (int i=0 ; i<emails.length ; i++) {
            Optional<User> exist = underTest.findByEmail(emails[i]);
            if (i != emails.length-1) {
                assertThat(exist).isNotEqualTo(Optional.empty());
                continue;
            }
            assertThat(exist).isEqualTo(Optional.empty());
        }
    }
}