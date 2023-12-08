defmodule Day1 do
  def part_1() do
    number_list_from_file("input.txt")
    |> Enum.sum()
  end

  def part_2() do
    freqs =
      MapSet.new()
      |> MapSet.put(0)

    number_list = number_list_from_file("input.txt") |> Enum.to_list()
    find_duplicate(number_list, freqs, 0)
  end

  defp find_duplicate(number_list, freqs, current_freq) do
    case helper(number_list, freqs, current_freq) do
      {:eof, freqs, current_freq} -> find_duplicate(number_list, freqs, current_freq)
      {:duplicate_found, current_freq} -> current_freq
      _ -> :error
    end
  end

  defp helper([], freqs, current_freq), do: {:eof, freqs, current_freq}

  defp helper([h | t], freqs, current_freq) do
    current_freq = current_freq + h

    if MapSet.member?(freqs, current_freq) do
      {:duplicate_found, current_freq}
    else
      freqs = MapSet.put(freqs, current_freq)
      helper(t, freqs, current_freq)
    end
  end

  defp number_list_from_file(path) do
    File.stream!(path)
    |> Stream.map(&String.replace(&1, "\n", ""))
    |> Stream.map(&String.to_integer/1)
  end
end
