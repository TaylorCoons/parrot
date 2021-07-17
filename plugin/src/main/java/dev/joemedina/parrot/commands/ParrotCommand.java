package dev.joemedina.parrot.commands;

import org.bukkit.command.Command;
import org.bukkit.command.CommandExecutor;
import org.bukkit.command.CommandSender;
import org.bukkit.entity.Player;

public class ParrotCommand implements CommandExecutor {

    @Override
    public boolean onCommand(CommandSender commandSender, Command command, String s, String[] args) {
        if(!(commandSender instanceof Player))
        {
            commandSender.sendMessage("This command can only be execute by a player!");
            return false;
        }

        Player player = (Player) commandSender;

        if(args.length > 0)
        {
            switch(args[1].toLowerCase())
            {
                case "list": ParrotCommand(player, args);
                case "help": HelpCommand(player);
                default: HelpCommand(player);
            }
        }
        return false;
    }

    private void HelpCommand(Player player)
    {
    }

    private void ParrotCommand(Player player, String[] args)
    {

    }
}
