import 'package:flutter/material.dart';
import 'package:frontend/src/api/recipes.dart';
import 'package:flutter_slidable/flutter_slidable.dart';
import 'package:frontend/src/my_recipes_feature/meals.dart';
import 'package:frontend/src/my_recipes_feature/recipes_provider.dart';
import 'package:frontend/src/my_recipes_feature/units.dart';
import 'package:frontend/src/search_feature/search_view.dart';
import 'package:provider/provider.dart';

typedef RefetchRecipe = void Function();

class SearchViewArguments {
  final int recipeId;
  final RefetchRecipe callback;

  SearchViewArguments(this.recipeId, this.callback);
}

class RecipeIdArguments {
  final int recipeId;

  RecipeIdArguments(this.recipeId);
}

class _MyRecipesState extends State<RecipeDetailsView> {
  @override
  void initState() {
    super.initState();
    WidgetsBinding.instance.addPostFrameCallback((_) {
      final recipesProvider =
          Provider.of<RecipesProvider>(context, listen: false);
      recipesProvider.getRecipe(widget.recipeId);
    });
  }

  @override
  Widget build(BuildContext context) {
    return Consumer<RecipesProvider>(
        builder: (context, recipesProvider, child) {
      return Scaffold(
        appBar: AppBar(title: const Text("Mein Rezept")),
        floatingActionButtonLocation: FloatingActionButtonLocation.endTop,
        floatingActionButton: FloatingActionButton(
          backgroundColor: Colors.green,
          onPressed: () => Navigator.pushNamed(context, SearchView.routeName,
              arguments: SearchViewArguments(widget.recipeId,
                  () => recipesProvider.getRecipe(widget.recipeId))),
          child: const Icon(Icons.add),
        ),
        body: Center(
          child: FutureBuilder<FullRecipe>(
            future: recipesProvider.futureRecipe,
            builder: (context, snapshot) {
              if (snapshot.hasData && recipesProvider.futureRecipe != null) {
                return Center(
                  child: Column(children: [
                    Text(snapshot.data!.name),
                    Text(
                        "${snapshot.data!.defaultPortions.toString()} Portionen"),
                    Text(meals[snapshot.data!.defaultMeal] ?? ""),
                    TextButton(
                        onPressed: () async {
                          final wasRemoved = await Provider.of<RecipesProvider>(
                                  context,
                                  listen: false)
                              .removeRecipe(widget.recipeId);
                          if (wasRemoved && context.mounted) {
                            Navigator.pop(context);
                          }
                        },
                        child: const Text("Löschen")),
                    Expanded(
                      child: ListView.builder(
                          itemCount: snapshot.data!.recipeIngredients.length,
                          itemBuilder: (BuildContext context, int index) {
                            return Slidable(
                              // Specify a key if the Slidable is dismissible.
                              key: const ValueKey(0),

                              // The start action pane is the one at the left or the top side.
                              startActionPane: ActionPane(
                                // A motion is a widget used to control how the pane animates.
                                motion: const ScrollMotion(),

                                // A pane can dismiss the Slidable.
                                dismissible: null,
                                // All actions are defined in the children parameter.
                                children: [
                                  SlidableAction(
                                    backgroundColor: const Color(0xFF21B7CA),
                                    foregroundColor: Colors.white,
                                    icon: Icons.share,
                                    label: 'Share',
                                    onPressed: (BuildContext context) {
                                      showModalBottomSheet<void>(
                                        context: context,
                                        builder: (BuildContext context) {
                                          return Container(
                                            height: 200,
                                            color: Colors.amber,
                                            child: Center(
                                              child: Column(
                                                mainAxisAlignment:
                                                    MainAxisAlignment.center,
                                                mainAxisSize: MainAxisSize.min,
                                                children: [
                                                  Text(snapshot
                                                      .data!
                                                      .recipeIngredients[index]
                                                      .name),
                                                  Padding(
                                                      padding: const EdgeInsets
                                                          .symmetric(
                                                          horizontal: 30.0,
                                                          vertical: 5.0),
                                                      child: Row(
                                                        mainAxisAlignment:
                                                            MainAxisAlignment
                                                                .spaceBetween,
                                                        children: [
                                                          Expanded(
                                                              child:
                                                                  TextFormField(
                                                            initialValue: snapshot
                                                                .data!
                                                                .recipeIngredients[
                                                                    index]
                                                                .amountPerPortion
                                                                .toString(),
                                                            decoration:
                                                                const InputDecoration(
                                                                    border:
                                                                        OutlineInputBorder()),
                                                          )),
                                                          const SizedBox(
                                                              width: 50),
                                                          Expanded(
                                                            child: TextFormField(
                                                                initialValue:
                                                                    units[snapshot
                                                                        .data!
                                                                        .recipeIngredients[
                                                                            index]
                                                                        .unit],
                                                                decoration: const InputDecoration(
                                                                    labelText:
                                                                        "Pro Portion",
                                                                    floatingLabelBehavior:
                                                                        FloatingLabelBehavior
                                                                            .always,
                                                                    border:
                                                                        OutlineInputBorder())),
                                                          )
                                                        ],
                                                      )),
                                                  ElevatedButton(
                                                    child: const Text(
                                                        'Close BottomSheet'),
                                                    onPressed: () =>
                                                        Navigator.pop(context),
                                                  ),
                                                ],
                                              ),
                                            ),
                                          );
                                        },
                                      );
                                    },
                                  ),
                                ],
                              ),
                              endActionPane: ActionPane(
                                motion: const ScrollMotion(),
                                dismissible:
                                    DismissiblePane(onDismissed: () async {
                                  await deleteRecipeIngredient(snapshot
                                      .data!.recipeIngredients[index].id);
                                }),
                                children: [
                                  SlidableAction(
                                    key: ValueKey(snapshot
                                        .data!.recipeIngredients[index].id),
                                    backgroundColor: const Color(0xFFFE4A49),
                                    foregroundColor: Colors.white,
                                    icon: Icons.delete,
                                    label: 'Delete',
                                    onPressed: (BuildContext context) async {
                                      await deleteRecipeIngredient(snapshot
                                          .data!.recipeIngredients[index].id);
                                      recipesProvider
                                          .getRecipe(widget.recipeId);
                                    },
                                  )
                                ],
                              ),

                              // The child of the Slidable is what the user sees when the
                              // component is not dragged.
                              child: ListTile(
                                  title: Text(snapshot
                                      .data!.recipeIngredients[index].name),
                                  subtitle: Row(
                                      mainAxisAlignment:
                                          MainAxisAlignment.spaceBetween,
                                      children: [
                                        Text(
                                            "${snapshot.data!.recipeIngredients[index].amountPerPortion.toString()} ${units[snapshot.data!.recipeIngredients[index].unit]}"),
                                        Text(snapshot.data!
                                            .recipeIngredients[index].market)
                                      ])),
                            );
                          }),
                    )
                  ]),
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
    });
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
