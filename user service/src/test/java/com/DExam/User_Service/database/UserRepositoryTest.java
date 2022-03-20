package com.DExam.User_Service.database;

import com.DExam.User_Service.model.User;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.orm.jpa.DataJpaTest;

import java.util.Optional;

import static org.assertj.core.api.AssertionsForClassTypes.assertThat;

@DataJpaTest
class UserRepositoryTest {
    @Autowired
    private UserRepository underTest;

    @Test
    void findByEmailTest() {
        String[] emails = new String[]{"eyad@gmail.com", "marwan@gmail.com",
                "ammar@gmail.com", "bodda@gmail.com"};
        for (int i = 0; i < emails.length - 1; i++) {
            User user = new User();
            user.setEmail(emails[i]);
            underTest.save(user);
        }
        for (int i = 0; i < emails.length; i++) {
            Optional<User> exist = underTest.findByEmail(emails[i]);
            if (i != emails.length - 1) {
                assertThat(exist).isNotEqualTo(Optional.empty());
                continue;
            }
            assertThat(exist).isEqualTo(Optional.empty());
        }
    }

    @Test
    void findByNationalIDTest() {
        String[] ids = new String[]{"0132156", "9784654",
                "9876541", "16531320231"};
        for (int i = 0; i < ids.length - 1; i++) {
            User user = new User();
            user.setNationalID(ids[i]);
            underTest.save(user);
        }

        for (int i = 0; i < ids.length; i++) {
            Optional<User> exist = underTest.findByNationalID(ids[i]);
            if (i != ids.length - 1) {
                assertThat(exist).isNotEqualTo(Optional.empty());
                continue;
            }
            assertThat(exist).isEqualTo(Optional.empty());
        }
    }
}