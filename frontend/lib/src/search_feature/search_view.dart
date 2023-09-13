import 'package:flutter/material.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';

class _SearchViewState extends State<SearchView> {
  @override
  void dispose() {
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(title: const Text("Nahrungsmittel")),
        body: Center(
            child: Column(children: [
          Text(widget.recipeId.toString()),
          Text(dotenv.env["APPLICATION_ID"]!)
        ])));
  }
}

class SearchView extends StatefulWidget {
  final int recipeId;
  const SearchView({super.key, required this.recipeId});

  @override
  State<SearchView> createState() => _SearchViewState();

  static const routeName = "/search";
}
