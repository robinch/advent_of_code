defmodule Day03 do
  def part_01() do
    input = File.read!("input") |> String.split("\n", trim: true)

    res =
      input
      |> Enum.map(&parse_valid_instructions/1)
      |> List.flatten()
      |> run_instructions()
      |> Enum.sum()

    IO.puts("Day 3 - Part 1: #{res}")
  end

  def part_02() do
    input = File.read!("input") |> String.split("\n", trim: true)

    res =
      input
      |> Enum.map(&parse_valid_instructions/1)
      |> List.flatten()
      |> run_extended_instructions()
      |> Enum.sum()

    IO.puts("Day 3 - Part 2: #{res}")
  end

  defp parse_valid_instructions(input_row) do
    # This pattern matches
    # - mul(n,m) where n and m are 1-3 digit numbers
    # - don't()
    # - do()
    pattern = ~r/(mul)\((\d{1,3}),(\d{1,3})\)|(don't)\(\)|(do)\(\)/

    pattern
    |> Regex.scan(input_row, capture: :all_but_first)
    |> Enum.map(fn
      [instruction, arg1, arg2] ->
        {String.to_existing_atom(instruction), String.to_integer(arg1), String.to_integer(arg2)}

      [_, _, _, "don't"] ->
        :do_not

      [_, _, _, _, "do"] ->
        :do
    end)
  end

  defp run_instructions(instructions) do
    instructions
    |> Enum.reduce([], fn
      {instruction, arg1, arg2}, acc -> [exec({instruction, arg1, arg2}) | acc]
      _, acc -> acc
    end)
    |> Enum.reverse()
  end

  defp run_extended_instructions(instructions) do
    {_enabled, executed_instruction} =
      instructions
      |> Enum.reduce({_enabled = true, []}, fn
        {instruction, arg1, arg2}, {true, executed_instructions} ->
          {true, [exec({instruction, arg1, arg2}) | executed_instructions]}

        :do_not, {_enabled, executed_instruction} ->
          {false, executed_instruction}

        :do, {_enabled, executed_instruction} ->
          {true, executed_instruction}

          _, acc -> acc 
      end)

    Enum.reverse(executed_instruction)
  end

  defp exec({:mul, arg1, arg2}) do
    arg1 * arg2
  end
end

Day03.part_01()
Day03.part_02()
