package com.DExam.User_Service.config;

import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.SignatureAlgorithm;
import org.springframework.stereotype.Service;

import java.util.Date;

@Service
public class JwtManager {
    private static final String SECRET_KEY = System.getenv("AUTH_SECRET_KEY");
    static byte[] secretKey = SECRET_KEY.getBytes();
    public static String generateToken(String email, String role) {
        Claims claims = Jwts.claims().setSubject(email);
        claims.put("UserRole", role);
        long expTime = System.currentTimeMillis() + 1000 * 60 * 60 * 10;
        Date exp = new Date(expTime);
        return Jwts.builder().setClaims(claims).setIssuedAt(new Date(System.currentTimeMillis())).setExpiration(exp)
                .signWith(SignatureAlgorithm.HS256, secretKey).compact();
    }
}
