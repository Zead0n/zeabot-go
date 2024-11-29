package response

import (
	"fmt"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/json"
)

func Update(content string) discord.MessageUpdate {
	return discord.MessageUpdate{
		Embeds: &[]discord.Embed{
			{
				Description: content,
				Color:       PrimaryColor,
			},
		},
	}
}

func Updatef(format string, a ...any) discord.MessageUpdate {
	return Update(fmt.Sprintf(format, a...))
}

func UpdateSuccess(content string) discord.MessageUpdate {
	return discord.MessageUpdate{
		Embeds: &[]discord.Embed{
			{
				Description: content,
				Color:       SuccessColor,
			},
		},
	}
}

func UpdateSuccessf(format string, a ...any) discord.MessageUpdate {
	return UpdateSuccess(fmt.Sprintf(format, a...))
}

func UpdateWarn(content string) discord.MessageUpdate {
	return discord.MessageUpdate{
		Embeds: &[]discord.Embed{
			{
				Description: content,
				Color:       WarnColor,
			},
		},
	}
}

func UpdateWarnf(format string, a ...any) discord.MessageUpdate {
	return UpdateWarn(fmt.Sprintf(format, a...))
}

func UpdateErr(message string, err error, a ...any) discord.MessageUpdate {
	return UpdateError(message + ": " + err.Error())
}

func UpdateError(message string, a ...any) discord.MessageUpdate {
	return discord.MessageUpdate{
		Embeds: &[]discord.Embed{
			{
				Description: fmt.Sprintf(message, a...),
				Color:       ErrorColor,
			},
		},
		Flags: json.Ptr(discord.MessageFlagEphemeral),
	}
}
