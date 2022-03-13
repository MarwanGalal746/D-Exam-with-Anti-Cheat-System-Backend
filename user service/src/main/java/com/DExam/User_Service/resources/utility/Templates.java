package com.DExam.User_Service.resources.utility;

import java.util.HashMap;
import java.util.Map;

public final class Templates {
    static final Map<String, String> templates = new HashMap<>();
    Templates(){
        templates.put("EMAIL VERIFICATION", "Verify your email to finish the process using the verification code: ");

    }
    public static final Map<String, String> getTemplates() {
        return templates;
    }
}
