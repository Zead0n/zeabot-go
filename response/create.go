package response

import (
	"fmt"

	"github.com/disgoorg/disgo/discord"
)

const (
	PrimaryColor = 0x4a90e2
	SuccessColor = 0x82dd55
	ErrorColor   = 0xe23636
	WarnColor    = 0xedb95e
)

func Create(content string) discord.MessageCreate {
	return discord.MessageCreate{
		Embeds: []discord.Embed{
			{
				Description: content,
				Color:       PrimaryColor,
			},
		},
	}
}

func Createf(format string, a ...any) discord.MessageCreate {
	return Create(fmt.Sprintf(format, a...))
}

func CreateSuccess(content string) discord.MessageCreate {
	return discord.MessageCreate{
		Embeds: []discord.Embed{
			{
				Description: content,
				Color:       SuccessColor,
			},
		},
		Flags: discord.MessageFlagEphemeral,
	}
}

func CreateSuccessf(format string, a ...any) discord.MessageCreate {
	return CreateSuccess(fmt.Sprintf(format, a...))
}

func CreateWarn(content string) discord.MessageCreate {
	return discord.MessageCreate{
		Embeds: []discord.Embed{
			{
				Description: content,
				Color:       WarnColor,
			},
		},
		Flags: discord.MessageFlagEphemeral,
	}
}

func CreateWarnf(format string, a ...any) discord.MessageCreate {
	return CreateWarn(fmt.Sprintf(format, a...))
}

func CreateErr(message string, err error, a ...any) discord.MessageCreate {
	return CreateError(message + ": " + err.Error())
}

func CreateError(message string, a ...any) discord.MessageCreate {
	return discord.MessageCreate{
		Embeds: []discord.Embed{
			{
				Description: fmt.Sprintf(message, a...),
				Color:       ErrorColor,
			},
		},
		Flags: discord.MessageFlagEphemeral,
	}
}
