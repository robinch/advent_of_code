defmodule AOC2021.Day10 do
  def part_01(path) do
    path
    |> input()
    |> error_score()
  end

  def part_02(path) do
    path
    |> input()
    |> completion_score()
  end

  defp error_score(navigation_subsystem) do
    Enum.reduce(navigation_subsystem, 0, fn line, acc ->
      case syntax_check(line) do
        {:error, {:illegal_character, c}} -> acc + illegal_character_score(c)
        _ -> acc
      end
    end)
  end

  defp illegal_character_score(")"), do: 3
  defp illegal_character_score("]"), do: 57
  defp illegal_character_score("}"), do: 1197
  defp illegal_character_score(">"), do: 25137

  defp completion_score(navigation_subsystem) do
    sorted_scores =
      Enum.reduce(navigation_subsystem, [], fn line, acc ->
        case syntax_check(line) do
          {:error, {:incomplete_line, characters}} ->
            [completion_score_for_incomplete_line(characters) | acc]

          _ ->
            acc
        end
      end)
      |> Enum.sort()

    Enum.at(sorted_scores, div(Enum.count(sorted_scores), 2))
  end

  defp completion_score_for_incomplete_line(incomplete_line) do
    Enum.map(incomplete_line, &closing_character/1)
    |> Enum.reduce(0, fn c, score ->
      score * 5 + completion_character_score(c)
    end)
  end

  defp completion_character_score(")"), do: 1
  defp completion_character_score("]"), do: 2
  defp completion_character_score("}"), do: 3
  defp completion_character_score(">"), do: 4

  defp syntax_check(line), do: syntax_check(line, [])

  defp syntax_check([], []), do: :ok
  defp syntax_check([], incomplete_line), do: {:error, {:incomplete_line, incomplete_line}}

  defp syntax_check([c | rest], opening_characters) when c in ["(", "[", "{", "<"] do
    syntax_check(rest, [c | opening_characters])
  end

  defp syntax_check([c | c_rest], [oc | oc_rest]) do
    cond do
      c == closing_character(oc) ->
        syntax_check(c_rest, oc_rest)

      :otherwise ->
        {:error, {:illegal_character, c}}
    end
  end

  defp closing_character("("), do: ")"
  defp closing_character("["), do: "]"
  defp closing_character("{"), do: "}"
  defp closing_character("<"), do: ">"

  defp input(path) do
    path
    |> File.read!()
    |> String.split("\n", trim: true)
    |> Enum.map(&String.codepoints/1)
  end
end
