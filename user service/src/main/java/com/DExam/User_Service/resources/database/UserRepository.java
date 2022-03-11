package com.DExam.User_Service.resources.database;

import com.DExam.User_Service.resources.modules.User;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;

@Repository
public interface UserRepository extends JpaRepository<User, Long> {
    @Query(value = "SELECT COUNT(*) FROM users WHERE email = ?1", nativeQuery = true)
    int isEmailExists(String email);
    @Query(value = "SELECT COUNT(*) FROM users WHERE nationalid = ?1", nativeQuery = true)
    int isNationalIdExists(String id);
}
