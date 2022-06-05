package com.DExam.User_Service.database;

import com.DExam.User_Service.domain.User;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Modifying;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;
import org.springframework.transaction.annotation.Transactional;

import java.util.Optional;

@Repository
public interface UserRepository extends JpaRepository<User, Long> {
    Optional <User> findByEmail(String email);

    Optional <User> findByNationalID(String nationalID);

    @Modifying
    @Transactional
    @Query(value = "UPDATE users SET password= ?2 WHERE email = ?1", nativeQuery = true)
    void updatePassword(String email, String password);

}
