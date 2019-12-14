defmodule Aoc2019.Day4 do
  @spec part_1(integer) :: boolean
  def part_1(digit) when is_integer(digit) do
    digit_list = to_digit_list(digit)

    has_two_adjacent_matching_digits(digit_list) &&
      never_decreasing(digit_list)
  end

  @spec part_2(integer) :: boolean
  def part_2(digit) do
    digit_list = to_digit_list(digit)

    has_two_adjacent_matching_digits(digit_list) &&
      has_matching_digit_group_of_size(digit_list, 2) &&
      never_decreasing(digit_list)
  end

  defp has_two_adjacent_matching_digits([_]), do: false
  defp has_two_adjacent_matching_digits([a, a | _]), do: true

  defp has_two_adjacent_matching_digits([_, a | rest]),
    do: has_two_adjacent_matching_digits([a | rest])

  defp has_matching_digit_group_of_size(digit_list, size) do
    {counter_list, current_counter, _prev_digit} =
      Enum.reduce(digit_list, {[], 1, nil}, fn
        digit, {counter_list, current_counter, prev_digit} when digit == prev_digit ->
          {counter_list, current_counter + 1, digit}

        digit, {counter_list, current_counter, _prev_digit} ->
          {[current_counter | counter_list], 1, digit}
      end)

    [current_counter | counter_list]
    |> Enum.member?(size)
  end

  defp never_decreasing([_]), do: true
  defp never_decreasing([a, b | _rest]) when a > b, do: false
  defp never_decreasing([_a, b | rest]), do: never_decreasing([b | rest])

  defp to_digit_list(digits) do
    digits
    |> Integer.to_string()
    |> String.graphemes()
    |> Enum.map(&String.to_integer/1)
  end
end
