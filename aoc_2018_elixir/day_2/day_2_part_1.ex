defmodule Day2.Part1 do
  def solve() do
    {total_doubles, total_triples} =
      read_file("input.txt")
      |> Enum.map(&list_to_count_map/1)
      |> Enum.map(&get_doubles_and_triples/1)
      |> sum_doubles_and_triples()

    total_doubles * total_triples
  end

  defp list_to_count_map(box_id) do
    Enum.reduce(String.codepoints(box_id), %{}, fn id_part, count_map ->
      Map.update(count_map, id_part, 1, &(&1 + 1))
    end)
  end

  defp get_doubles_and_triples(count_map) do
    Enum.reduce(count_map, {0, 0}, fn
      {_key, 2}, {_double, triple} -> {1, triple}
      {_key, 3}, {double, _triple} -> {double, 1}
      _, acc -> acc
    end)
  end

  defp sum_doubles_and_triples(doubles_and_triples) do
    doubles_and_triples
    |> Enum.reduce({0, 0}, fn
      {double, triple}, {total_doubles, total_triples} ->
        {total_doubles + double, total_triples + triple}
    end)
  end

  defp read_file(path) do
    File.stream!(path)
    |> Stream.map(&String.replace(&1, "\n", ""))
  end
end
