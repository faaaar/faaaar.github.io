jQuery(document).ready(function ($) {
  var word_idx = 0;
  var words = [
    "富强(Prosperity)",
    "民主(Democracy)",
    "文明(Civility)",
    "和谐(Harmony)",
    "自由(Freedom)",
    "平等(Equality)",
    "公正(Justice)",
    "法治(Rule of Law)",
    "爱国(Patriotism)",
    "敬业(Dedication)",
    "诚信(Integrity)",
    "友善(Friendship)",
  ];

  $("body").click(function (e) {
    var $i = $("<span style=\"user-select: none;\" />").text(words[word_idx]);
    var x = e.pageX;
    var y = e.pageY;

    word_idx = (word_idx + 1) % words.length;
    $i.css({
      "z-index": 9999,
      "top": y - 20,
      "left": x,
      "position": "absolute",
      "font-weight": "bold",
      "color": "#ff6651"
    });
    $("body").append($i);
    $i.animate({
      "top": y - 180,
      "opacity": 0
    }, 1500, function () {
      $i.remove();
    });
  });
});
