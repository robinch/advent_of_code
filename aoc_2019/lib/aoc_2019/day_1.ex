defmodule Aoc2019.Day1 do
  # Fuel required to launch a given module is based on its mass.
  # Specifically, to find the fuel required for a module,
  # take its mass, divide by three, round down, and subtract 2.
  @spec part_1(integer) :: integer
  def part_1(mass) when is_integer(mass) do
    calculate_fuel(mass)
  end

  @spec part_1([integer]) :: integer
  def part_1(masses) when is_list(masses) do
    masses
    |> Enum.map(&calculate_fuel/1)
    |> Enum.sum()
  end

  # Fuel itself requires fuel just like a module - take its mass, divide by three,
  # round down, and subtract 2. However, that fuel also requires fuel, and that
  # fuel requires fuel, and so on. Any mass that would require negative fuel should
  # instead be treated as if it requires zero fuel; the remaining mass, if any, is
  # instead handled by wishing really hard, which has no mass and is outside the
  # scope of this calculation.
  @spec part_2(integer) :: integer
  def part_2(mass) when is_integer(mass) do
    recursively_calculate_fuel(mass, 0)
  end

  @spec part_2([integer]) :: integer
  def part_2(masses) when is_list(masses) do
    masses
    |> Enum.map(&recursively_calculate_fuel(&1, 0))
    |> Enum.sum()
  end

  defp recursively_calculate_fuel(mass, acc) do
    fuel = calculate_fuel(mass)

    cond do
      fuel > 0 -> recursively_calculate_fuel(fuel, acc + fuel)
      true -> acc
    end
  end

  defp calculate_fuel(mass) do
    div(mass, 3) - 2
  end
end
