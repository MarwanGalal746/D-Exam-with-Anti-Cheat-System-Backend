package com.DExam.User_Service.resources.database;

import com.DExam.User_Service.resources.entity.User;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.Optional;

@Repository
public interface UserRepository extends JpaRepository<User, Long> {
    Optional <User> findByEmail(String email);

    Optional <User> findByNationalID(String nationalID);
}
