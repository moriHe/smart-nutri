import 'package:flutter/foundation.dart';
import 'package:frontend/src/api/recipes.dart';

class RecipesProvider with ChangeNotifier, DiagnosticableTreeMixin {
  late Future<List<ShallowRecipe>> _futureRecipes;
  Future<FullRecipe>? _futureRecipe;

  void getRecipes() {
    _futureRecipes = fetchRecipes();
    notifyListeners();
  }

  void getRecipe(int id) {
    _futureRecipe = null;
    _futureRecipe = fetchRecipe(id);
    notifyListeners();
  }

  Future<bool> removeRecipe(int id) async {
    int responseCode = await deleteRecipe(id);
    if (responseCode == 200) {
      getRecipes();
      return true;
    }
    if (responseCode == 200) {
      notifyListeners();
      return true;
    }
    return false;
  }

  Future<List<ShallowRecipe>> get futureRecipes => _futureRecipes;
  Future<FullRecipe>? get futureRecipe => _futureRecipe;

  /// Makes `Counter` readable inside the devtools by listing all of its properties
  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties.add(DiagnosticsProperty<Future<List<ShallowRecipe>>>(
        "futureRecipes", futureRecipes));
    properties.add(
        DiagnosticsProperty<Future<FullRecipe>?>("futureRecipe", futureRecipe));
  }
}
