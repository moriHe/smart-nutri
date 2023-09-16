import 'package:flutter/foundation.dart';
import 'package:frontend/src/api/recipes.dart';

class RecipesProvider with ChangeNotifier, DiagnosticableTreeMixin {
  late Future<List<ShallowRecipe>> _futureRecipes;

  void getRecipes() {
    _futureRecipes = fetchRecipes();
    notifyListeners();
  }

  Future<List<ShallowRecipe>> get futureRecipes => _futureRecipes;

  /// Makes `Counter` readable inside the devtools by listing all of its properties
  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties.add(DiagnosticsProperty<Future<List<ShallowRecipe>>>(
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
