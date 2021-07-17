package dev.joemedina.parrot;

import dev.joemedina.parrot.commands.ParrotCommand;
import org.bukkit.plugin.java.JavaPlugin;

public class Parrot extends JavaPlugin {
    @Override
    public void onEnable() {
        super.onEnable();
        getCommand("parrot").setExecutor(new ParrotCommand());
    }

    @Override
    public void onDisable() {
        super.onDisable();
    }
}
