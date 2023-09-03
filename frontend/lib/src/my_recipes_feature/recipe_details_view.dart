import 'package:flutter/material.dart';
import 'package:frontend/src/api/recipes.dart';

class _MyRecipesState extends State<RecipeDetailsView> {
  late Future<FullRecipe> futureRecipe;

  @override
  void initState() {
    super.initState();
    futureRecipe = fetchRecipe(1);
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text("Meine Rezepte")),
      body: Center(
        child: FutureBuilder<FullRecipe>(
          future: futureRecipe,
          builder: (context, snapshot) {
            if (snapshot.hasData) {
              return Row(
                children: [
                  Text(snapshot.data!.name),
                  Text(snapshot.data!.defaultMeal)
                ],
              );
            } else if (snapshot.hasError) {
              return Text('${snapshot.error}');
            }

            // By default, show a loading spinner.
            return const CircularProgressIndicator();
          },
        ),
      ),
    );
  }
}

/// Displays detailed information about a SampleItem.
class RecipeDetailsView extends StatefulWidget {
  const RecipeDetailsView({super.key});

  @override
  State<RecipeDetailsView> createState() => _MyRecipesState();

  static const routeName = '/recipe';
}
