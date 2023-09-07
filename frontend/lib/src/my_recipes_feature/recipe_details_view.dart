import 'package:flutter/material.dart';
import 'package:frontend/src/api/recipes.dart';

class RecipeDetailsViewArguments {
  final int recipeId;

  RecipeDetailsViewArguments(this.recipeId);
}

class _MyRecipesState extends State<RecipeDetailsView> {
  late Future<FullRecipe> futureRecipe;
  @override
  void initState() {
    super.initState();
    futureRecipe = fetchRecipe(widget.recipeId);
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text("Mein Rezept")),
      body: Center(
        child: FutureBuilder<FullRecipe>(
          future: futureRecipe,
          builder: (context, snapshot) {
            if (snapshot.hasData) {
              return Center(
                child: Column(children: [
                  Text(snapshot.data!.name),
                  Text(
                      "${snapshot.data!.defaultPortions.toString()} Portionen"),
                  Text(snapshot.data!.defaultMeal),
                  Expanded(
                    child: ListView.builder(
                        itemCount: snapshot.data!.recipeIngredients.length,
                        itemBuilder: (BuildContext context, int index) {
                          return Card(
                              margin: const EdgeInsets.only(
                                  right: 20.0,
                                  left: 20.0,
                                  top: 10.0,
                                  bottom: 10.0),
                              child: ListTile(
                                  title: Text(snapshot
                                      .data!.recipeIngredients[index].name),
                                  subtitle: Row(
                                      mainAxisAlignment:
                                          MainAxisAlignment.spaceBetween,
                                      children: [
                                        Text(
                                            "${snapshot.data!.recipeIngredients[index].amountPerPortion.toString()} ${snapshot.data!.recipeIngredients[index].unit}"),
                                        Text(snapshot.data!
                                            .recipeIngredients[index].market)
                                      ])));
                        }),
                  )
                ]),
              );
              // return Row(
              //   children: [
              //     Text(snapshot.data!.name),
              //     Text(snapshot.data!.defaultMeal)
              //   ],
              // );
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
  final int recipeId;
  const RecipeDetailsView({super.key, required this.recipeId});

  @override
  State<RecipeDetailsView> createState() => _MyRecipesState();

  static const routeName = '/recipe';
}
