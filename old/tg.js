//const TelegramBot = require('node-telegram-bot-api');

const Telegram = require('telegraf/telegram')



const token = '787991996:AAEC0_MHueKV4bETpzZb3ZRfPa2wqfgATX0';

// Create a bot that uses 'polling' to fetch new updates
//
	const bot = new Telegram(token)
	bot.sendMessage('@gongo_hezner', 'test')
