package tpl

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/anihouse/bot/config"
	"github.com/bwmarrin/discordgo"
)

var (
	tpls *template.Template
)

func Init() {
	tpls = template.New("").Funcs(funcs)

	fmt.Println("Loading templates:")
	err := filepath.Walk(config.Bot.Templates, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		name, err := filepath.Rel(config.Bot.Templates, path)
		if err != nil {
			return err
		}

		_, err = tpls.New(name).Parse(string(data))
		return err
	})

	if err != nil {
		log.Fatal(err)
	}

	for _, tpl := range tpls.Templates() {
		fmt.Println("Load", tpl.Name())
	}
}

func get(name string, data interface{}) (*schema, error) {
	buf := bytes.NewBufferString("")
	err := tpls.ExecuteTemplate(buf, name, data)
	if err != nil {
		return nil, err
	}

	var result schema
	s := bytes.NewBuffer(normalizeSpaces(buf.Bytes()))
	if err := xml.NewDecoder(s).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func ToSend(name string, data interface{}) (*discordgo.MessageSend, error) {
	m, err := get(name, data)
	if err != nil {
		return nil, err
	}

	result := new(discordgo.MessageSend)

	if content := strings.TrimSpace(m.Content); content != "" {
		result.Content = content
	}

	if embed := m.Embed; embed != nil {
		result.Embed = new(discordgo.MessageEmbed)

		if title := embed.Title; title != nil {
			result.Embed.Title = *title
		}

		if color := embed.Color; color != nil {
			if len([]rune(*color)) != 7 {
				return nil, fmt.Errorf("Unexpected color format: '%s'", *color)
			}

			c, err := strconv.ParseInt((*color)[1:], 16, 32)
			if err != nil {
				return nil, err
			}
			result.Embed.Color = int(c)
		}

		if description := strings.TrimSpace(embed.Description); description != "" {
			result.Embed.Description = description
		}

		if footer := embed.Footer; footer != nil {
			result.Embed.Footer = new(discordgo.MessageEmbedFooter)
			result.Embed.Footer.Text = *footer
		}

		if fields := embed.Fields; fields != nil {
			result.Embed.Fields = make([]*discordgo.MessageEmbedField, 0)
			for _, field := range *fields {
				result.Embed.Fields = append(result.Embed.Fields, &discordgo.MessageEmbedField{
					Inline: field.Inline,
					Name:   strings.TrimSpace(field.Name),
					Value:  strings.TrimSpace(field.Value),
				})
			}
		}
	}

	return result, nil
}

func ToEdit(name string, data interface{}) (*discordgo.MessageEdit, error) {
	m, err := get(name, data)
	if err != nil {
		return nil, err
	}

	result := new(discordgo.MessageEdit)

	if content := strings.TrimSpace(m.Content); content != "" {
		result.Content = &content
	}

	if embed := m.Embed; embed != nil {
		result.Embed = new(discordgo.MessageEmbed)

		if title := embed.Title; title != nil {
			result.Embed.Title = *title
		}

		if color := embed.Color; color != nil {
			if len([]rune(*color)) != 7 {
				return nil, fmt.Errorf("Unexpected color format: '%s'", *color)
			}

			c, err := strconv.ParseInt((*color)[1:], 16, 32)
			if err != nil {
				return nil, err
			}
			result.Embed.Color = int(c)
		}

		if description := strings.TrimSpace(embed.Description); description != "" {
			result.Embed.Description = description
		}

		if footer := embed.Footer; footer != nil {
			result.Embed.Footer = new(discordgo.MessageEmbedFooter)
			result.Embed.Footer.Text = *footer
		}

		if fields := embed.Fields; fields != nil {
			result.Embed.Fields = make([]*discordgo.MessageEmbedField, 0)
			for _, field := range *fields {
				result.Embed.Fields = append(result.Embed.Fields, &discordgo.MessageEmbedField{
					Inline: field.Inline,
					Name:   strings.TrimSpace(field.Name),
					Value:  strings.TrimSpace(field.Value),
				})
			}
		}
	}
	return result, nil
}
