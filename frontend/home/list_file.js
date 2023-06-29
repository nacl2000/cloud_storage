$(function() {
    $.ajax({
        url: 'file-list',
        type: 'GET',
        dataType: 'json',
        success: function(data) {
            $.each(data, function(index, file) {
                var li = $('<li>');
                var a = $('<a>').attr('href', file.url);
                var img = $('<img>').attr('src', file.icon);
                var span = $('<span>').text(file.name);
                a.append(img).append(span);
                li.append(a);
                $('#file-list').append(li);
            });
        }
    });
});
