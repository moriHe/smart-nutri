import 'package:flutter/foundation.dart';
import 'package:frontend/src/api/recipes.dart';

class RecipesProvider with ChangeNotifier, DiagnosticableTreeMixin {
  late List<ShallowRecipe> _futureRecipes;

  void getRecipes() async {
    _futureRecipes = await fetchRecipes();
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

  List<ShallowRecipe> get futureRecipes => _futureRecipes;

  /// Makes `Counter` readable inside the devtools by listing all of its properties
  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties.add(DiagnosticsProperty<List<ShallowRecipe>>(
        "futureRecipes", futureRecipes));
  }
}

class RecipeProvider with ChangeNotifier, DiagnosticableTreeMixin {
  late Future<FullRecipe> _futureRecipe;

  void getRecipe(int id) {
    _futureRecipe = fetchRecipe(id);
    notifyListeners();
  }

  Future<FullRecipe> get futureRecipe => _futureRecipe;

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties.add(
        DiagnosticsProperty<Future<FullRecipe>>("futureRecipes", futureRecipe));
  }
}
