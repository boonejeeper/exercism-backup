public class Lasagna {

    private static final int EXPECTED_MINUTES_IN_OVEN = 40;
    private static final int PREPARATION_TIME_PER_LAYER = 2;
    
    public int expectedMinutesInOven() {
        return 40;
    }

    public int remainingMinutesInOven(int timeSpentInOven) {
        return expectedMinutesInOven() - timeSpentInOven;
    }

    public int preparationTimeInMinutes(int numberOfLayers) {
        return numberOfLayers * 2;
    }
    
    public int totalTimeInMinutes(int numberOfLayers, int timeSpentInOven) {
        return preparationTimeInMinutes(numberOfLayers) + timeSpentInOven;
    }
}
