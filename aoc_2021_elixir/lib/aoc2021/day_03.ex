defmodule AOC2021.Day03 do
  def part_01(path) do
    gamma = path |> read_input() |> gamma()

    epsilon = Enum.map(gamma, &(1 - &1))

    to_integer(gamma) * to_integer(epsilon)
  end

  def part_02(path) do
    inputs = path |> read_input() |> to_indexed_input()
    o2 = inputs |> o2() |> to_integer()
    co2 = inputs |> co2() |> to_integer()

    o2 * co2
  end

  defp read_input(path) do
    path
    |> File.read!()
    |> String.split("\n", trim: true)
    |> Enum.map(fn bits ->
      bits
      |> String.graphemes()
      |> Enum.map(&String.to_integer/1)
    end)
  end

  defp gamma(inputs), do: most_frequent_bits(inputs, [])

  defp most_frequent_bits([[] | _], acc), do: acc |> Enum.reverse()

  defp most_frequent_bits(inputs, acc) do
    most_frequent_bit =
      inputs
      |> Enum.map(fn [h | _] -> h end)
      |> most_frequent_value()

    rests = Enum.map(inputs, &tl/1)

    most_frequent_bits(rests, [most_frequent_bit | acc])
  end

  defp o2(input) do
    index = index_that_meets_bit_criteria(input, &most_frequent_value/1)
    input[index]
  end

  defp co2(input) do
    index = index_that_meets_bit_criteria(input, &least_frequent_value/1)
    input[index]
  end

  defp index_that_meets_bit_criteria([{index, _bits}], _), do: index

  defp index_that_meets_bit_criteria(inputs, frequency_function) do
    pivot =
      inputs
      |> Enum.map(fn {_index, [h | _]} -> h end)
      |> frequency_function.()

    new_inputs =
      Enum.reduce(inputs, [], fn
        {index, [^pivot | rest]}, acc -> [{index, rest} | acc]
        _, acc -> acc
      end)

    index_that_meets_bit_criteria(new_inputs, frequency_function)
  end

  defp to_indexed_input(input) do
    input
    |> Enum.with_index()
    |> Map.new(fn {value, key} -> {key, value} end)
  end

  defp to_integer(list) do
    max_bit_pos = Enum.count(list) - 1

    list
    |> Enum.with_index()
    |> Enum.reduce(0, fn {bit_val, bit_pos}, acc ->
      acc + bit_val * :math.pow(2, max_bit_pos - bit_pos)
    end)
    |> trunc()
  end

  defp most_frequent_value(list) do
    freqs = Enum.frequencies(list)

    cond do
      freqs[1] >= freqs[0] -> 1
      :otherwise -> 0
    end
  end

  defp least_frequent_value(list) do
    1 - most_frequent_value(list)
  end
end
