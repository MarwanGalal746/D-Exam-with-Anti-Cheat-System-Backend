package dexam.desktopapplication.anticheat;

import dexam.desktopapplication.Main;
import dexam.desktopapplication.messaging.RabbitMq;

import java.io.IOException;
import java.util.ArrayList;

public class AntiCheat {
    public static void initiate(String browser) throws IOException, InterruptedException {

        ProcessManager processManager = new ProcessManager();

        ArrayList<String> p = new ArrayList<>();
        p.add("java.exe");
        p.add("javaw.exe");
        p.add(browser);
        p.add("discord.exe");

        processManager.addToIgnoreList(p);
        processManager.killProcesses();
        Thread newThread = new Thread(() -> {
            try {
                boolean result = processManager.monitorActivity();
                if(result) {
                    RabbitMq.send("close-"+Main.userId);
                    Main.stopApp();
                }
            } catch (IOException e) {
                throw new RuntimeException(e);
            } catch (InterruptedException e) {
                throw new RuntimeException(e);
            }
        });
        newThread.start();

        return;
    }
}
