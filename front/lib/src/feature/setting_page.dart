import 'package:flutter/material.dart';

/// Displays a list of SampleItems.
class SettingPage extends StatelessWidget {
  const SettingPage({super.key,});

  static const routeName = '/edit';

  @override
  Widget build(BuildContext context) {
    return Container(
      child: Text("this is setting page"),
    );
  }
}
