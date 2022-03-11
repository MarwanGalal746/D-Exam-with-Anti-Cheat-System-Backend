package com.DExam.User_Service.resources.database;

import com.DExam.User_Service.resources.modules.User;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface UserRepository extends JpaRepository<User, Long> {
    User findByEmail(String email);
  
    User findByNationalID(String nationalID);
}
