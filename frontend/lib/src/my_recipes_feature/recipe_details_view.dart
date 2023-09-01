import 'package:flutter/material.dart';

/// Displays detailed information about a SampleItem.
class RecipeDetailsView extends StatelessWidget {
  const RecipeDetailsView({super.key});

  static const routeName = '/recipe';

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Item Details'),
      ),
      body: const Center(
        child: Text('More Information Here'),
      ),
    );
  }
}
