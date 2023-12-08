defmodule Day5 do
  def solve_part_1() do
    File.stream!("input.txt")
    |> Stream.map(&String.to_charlist/1)
    |> Enum.to_list()
    |> hd()
    |> reacted_polymer()
    |> length()
  end

  def solve_part_2() do
    polymer =
      File.stream!("input.txt")
      |> Stream.map(&String.to_charlist/1)
      |> Enum.to_list()
      |> hd()

    Enum.map(?A..?Z, fn
      remove_type -> length(reacted_polymer(polymer, remove_type))
    end)
    |> Enum.min()
  end

  def reacted_polymer(polymer, remove_type \\ "") do
    Enum.reduce(polymer, _stack = [], fn
      unit, list when unit == remove_type or unit == remove_type + 32 ->
        list

      unit, [] ->
        [unit]

      unit, stack = [stored_unit | rest] ->
        if same_type_opposite_polarity?(stored_unit, unit) do
          rest
        else
          [unit | stack]
        end
    end)
  end

  def same_type_opposite_polarity?(unit1, unit2) do
    abs(unit1 - unit2) == 32
  end
end
