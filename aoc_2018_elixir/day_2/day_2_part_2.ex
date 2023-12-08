defmodule Day2.Part2 do
  def solve() do
    id_list =
      read_file("input.txt")
      |> Enum.to_list()

    id_length = id_list |> hd() |> String.length()

    # TODO: MAKE THIS PART PRETTIER
    removed_letter_id_matrix =
      0..(id_length - 1)
      |> Enum.map(fn n ->
        Enum.map(id_list, fn id ->
          remove_letter(id, n)
        end)
      end)

    removed_letter_id_matrix
    |> find_duplicates()
  end

  def find_duplicates([]), do: :no_prototype_exists

  def find_duplicates([h | t]) do
    case find_dups_helper(h, MapSet.new()) do
      :nothing_found -> find_duplicates(t)
      {:prototype_found, id} -> id
    end
  end

  def find_dups_helper([], _set), do: :nothing_found

  def find_dups_helper([id | ids], set) do
    if MapSet.member?(set, id) do
      {:prototype_found, id}
    else
      set = MapSet.put(set, id)
      find_dups_helper(ids, set)
    end
  end

  def remove_letter(string, index) do
    String.slice(string, 0, index) <> String.slice(string, index + 1, String.length(string))
  end

  defp read_file(path) do
    File.stream!(path)
    |> Stream.map(&String.replace(&1, "\n", ""))
  end
end
