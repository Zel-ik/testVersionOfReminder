package main

import (
	"errors"
	"fmt"

	"gopkg.in/telebot.v3"
)

// Middleware для проверки админов
func AdminMiddleware(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		admins, err := c.Bot().AdminsOf(c.Chat())
		if err != nil {
			return fmt.Errorf("ошибка при получении списка админов: %w", err)
		}

		isAdmin := false
		for _, admin := range admins {
			if admin.User.ID == c.Sender().ID {
				isAdmin = true
				break
			}
		}

		if !isAdmin {
			c.Bot().Send(c.Chat(), "Власть над ботом принадлежит лишь админу")
			return errors.New("команда доступна только администраторам")
		}

		return next(c)
	}
}
