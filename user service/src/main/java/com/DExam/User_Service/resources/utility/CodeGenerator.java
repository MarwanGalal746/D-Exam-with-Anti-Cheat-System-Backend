package com.DExam.User_Service.resources.utility;

import java.util.Random;

public final class CodeGenerator {

    public static String generateCode(){
        String code=generateCodeUtil(48,57, 2);
        code+=generateCodeUtil(65,90, 3);
        code+=generateCodeUtil(48,57, 1);
        return code;
    }

    private static String generateCodeUtil(int leftLimit, int rightLimit, int targetStringLength ) {
        Random random = new Random();
        return random.ints(leftLimit, rightLimit + 1)
                .limit(targetStringLength)
                .collect(StringBuilder::new, StringBuilder::appendCodePoint, StringBuilder::append)
                .toString();
    }

}
