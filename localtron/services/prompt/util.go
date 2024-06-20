/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 *
 * For commercial use, a separate license must be obtained by purchasing from The Authors.
 * For commercial licensing inquiries, please contact The Authors listed in the AUTHORS file.
 */
package promptservice

import (
	"strings"

	"github.com/singulatron/singulatron/localtron/clients/llm"
)

func llmResponseToText(responses []*llm.CompletionResponse) string {
	var result strings.Builder

	first := true
	for _, v := range responses {
		if len(v.Choices) == 0 {
			continue
		}
		choice := v.Choices[0]

		var textToAdd string
		if strings.Contains(result.String(), "```") {
			// Handling for inline code formatting if the resulting string is already within a code block
			count := strings.Count(result.String(), "```")
			if count%2 == 1 { // If the count of ``` is odd, we are inside a code block
				textToAdd = choice.Text // No escaping needed inside code block
			} else {
				textToAdd = escapeHtml(choice.Text) // Apply HTML escaping when outside code blocks
			}
		} else {
			textToAdd = escapeHtml(choice.Text) // Apply HTML escaping if there is no code block
		}

		if first {
			textToAdd = strings.TrimLeft(textToAdd, " ")
			first = false
		}

		result.WriteString(textToAdd)

		if choice.FinishReason == "stop" {
			break
		}
	}

	return result.String()
}

func escapeHtml(input string) string {
	replacer := strings.NewReplacer(
		"&", "&amp;",
		"<", "&lt;",
		">", "&gt;",
		"\"", "&quot;",
		"'", "&#39;",
	)
	return replacer.Replace(input)
}

func errToString(err error) string {
	if err != nil {
		return err.Error()
	}

	return ""
}
